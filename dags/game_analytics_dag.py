from airflow import DAG
from airflow.operators.bash import BashOperator
from airflow.utils.dates import days_ago
from datetime import timedelta

default_args = {
    'owner': 'airflow',
    'depends_on_past': False,
    'email_on_failure': False,
    'retries': 1,
    'retry_delay': timedelta(minutes=5),
}

with DAG(
    'game_analytics_full_cycle',
    default_args=default_args,
    description='Запуск цикла трансформации данных из Bronze в Gold',
    schedule_interval='@hourly', 
    start_date=days_ago(1),
    catchup=False,
) as dag:

    # Running dbt deps
    dbt_deps = BashOperator(
        task_id='dbt_deps',
        bash_command='cd /home/airflow/gcs/dbt && dbt deps'
    )

    # Running dbt run (Bronze -> Silver -> Gold)
    dbt_run = BashOperator(
        task_id='dbt_run',
        bash_command='cd /home/airflow/gcs/dbt && dbt run'
    )

    # Running dbt test (Data Quality)
    dbt_test = BashOperator(
        task_id='dbt_test',
        bash_command='cd /home/airflow/gcs/dbt && dbt test'
    )

    dbt_deps >> dbt_run >> dbt_test