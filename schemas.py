from pydantic import BaseModel

class Item(BaseModel):
    id: int
    value: str

class ItemCreate(BaseModel):
    value: str 