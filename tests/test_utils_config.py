import pytest

from src.utils.config import (
    DatabaseConfiguration,
    Configuration
)

TEST_ENV_VARIABLES = {
    "DATABASE_HOST": "localhost",
    "DATABASE_PORT": "3306",
    "DATABASE_NAME": "test",
    "DATABASE_USER": "test",
    "DATABASE_PASSWORD": "test",
    "DATABASE_TYPE": "mysql",
}


class TestDatabaseConfiguration:
    def test_instantiation(self):
        db_config = DatabaseConfiguration()
        assert db_config is not None

    def test_set_property(self):
        db_config = DatabaseConfiguration()
        db_config.set_property("host", "localhost")
        assert db_config.host == "localhost"


class TestConfiguration:
    def test_instantiation(self):
        config = Configuration()
        assert config is not None

    def test_empty_database_configuration(self):
        config = Configuration()
        config.load_from_env()
        assert config.database is not None

        assert config.database.host is None
        assert config.database.port is None
        assert config.database.name is None
        assert config.database.user is None
        assert config.database.password is None
        assert config.database.type is None
