from airflow.models import DagBag

def test_dag_bag_import():
    """Import Errors"""
    dag_bag = DagBag(dag_folder="dags/", include_examples=False)
    assert len(dag_bag.import_errors) == 0, f"Ошибки в DAG-ах: {dag_bag.import_errors}"

def test_dag_ids():
    """Checking for the presence of specific DAGs"""
    dag_bag = DagBag(dag_folder="dags/", include_examples=False)
    assert "game_analytics_full_cycle" in dag_bag.dag_ids