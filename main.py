from fastapi import FastAPI, HTTPException
from fastapi.middleware.cors import CORSMiddleware
from typing import List
try:
    from .schemas import Item, ItemCreate
    from .storage import get_items, create_item, update_item, delete_item
except ImportError:
    from schemas import Item, ItemCreate
    from storage import get_items, create_item, update_item, delete_item

app = FastAPI()

app.add_middleware(
    CORSMiddleware,
    allow_origins=["http://localhost:3000"],
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)

@app.get("/items", response_model=List[Item])
def read_items():
    return get_items()

@app.post("/items", response_model=Item)
def add_item(item: ItemCreate):
    return create_item(item)

@app.put("/items/{item_id}", response_model=Item)
def edit_item(item_id: int, item: ItemCreate):
    updated = update_item(item_id, item)
    if updated:
        return updated
    raise HTTPException(status_code=404, detail="Item not found")

@app.delete("/items/{item_id}")
def remove_item(item_id: int):
    deleted = delete_item(item_id)
    if deleted:
        return {"ok": True}
    raise HTTPException(status_code=404, detail="Item not found") 