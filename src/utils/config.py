import logging
import os

logging_format = "%(asctime)s - %(name)s - %(levelname)s - %(message)s"

logging.basicConfig(
    level=logging.INFO
    logging_format=logging_format,
    datefmt="%Y-%m-%d %H:%M:%S"
)

logger = logging.getLogger(__name__)


class DatabaseConfiguration:
    """
    Database configuration class.
    """
    __SUPPORTED_DATABASE_TYPES = [
        "mysql",
        "postgresql",
        "sqlite"
    ]

    def __init__(self) -> None:
        """
        Constructor for the DatabaseConfiguration class.
        """
        self.host = None
        self.port = None
        self.name = None
        self.user = None
        self.password = None
        self.type = None

    def set_property(self, name: str, value: str) -> None:
        """
        Set a property on the configuration object.
        """
        if not hasattr(self, name):
            raise AttributeError(f"Invalid property: {name}")

        if name == "type" and value not in self.__SUPPORTED_DATABASE_TYPES:
            raise ValueError(f"Invalid database type: {value}")

        setattr(self, name, value)

    @property
    def uri(self):
        return (
            f"{self.type}://{self.user}:"
            f"{self.password}@{self.host}:"
            f"{self.port}/{self.name}"
        )


class Configuration:
    """
    Configuration class for the application.
    """
    __DATABASE_ENV_VARIABLES = [
        "DATABASE_HOST",
        "DATABASE_PORT",
        "DATABASE_NAME",
        "DATABASE_USER",
        "DATABASE_PASSWORD",
        "DATABSE_TYPE",
    ]
    __LOGGING_ENV_VARIABLES = [
        "LOG_LEVEL",
        "LOG_FILE",
    ]

    def __init__(self) -> None:
        """
        Constructor for the Configuration class.
        """
        if os.getenv("LOAD_FROM_ENVIRONMENT"):
            self.load_from_env()

    def load_from_env(self) -> None:
        """
        Load the configuration from the environment.
        """
        self.database = DatabaseConfiguration()
        for env_var in self.__DATABASE_ENV_VARIABLES:
            if os.getenv(env_var):
                self.database.set_property(env_var.lower().strip("database_"),
                                           os.getenv(env_var))

        