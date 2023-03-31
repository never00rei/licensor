import logging
import os
import configparser


class Configuration:
    """
    Configuration class for the application.
    """
    DEFAULT_CONFIG_PATH="/etc/licensor"
    DEFAULT_CONFIG_FILE="licensor.conf"

    __EXPECTED_CONFIG_SECTIONS = [
        "database",
        "logging",
    ]

    __ENVIRONMENT_VARIABLES = [
        "DATABASE_HOST",
        "DATABASE_PORT",
        "DATABASE_NAME",
        "DATABASE_USER",
        "DATABASE_PASSWORD",
        "DATABSE_TYPE",
        "LOG_LEVEL",
        "LOAD_FROM_ENVIRONMENT",
    ]

    def __init__(self) -> None:
        """
        Constructor for the Configuration class.
        """
        pass

    def load_from_env(self) -> None:
        """
        Load the configuration from the environment.
        """
        pass
        