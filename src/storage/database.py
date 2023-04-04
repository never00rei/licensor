from sqlalchemy import create_engine
from sqlalchemy.orm import sessionmaker
from sqlalchemy.ext.declarative import declarative_base

from utils.config import Configuration

config = Configuration()

engine = create_engine(
    config.database.uri
)

Session = sessionmaker(bind=engine)

base = declarative_base()
