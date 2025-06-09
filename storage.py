from typing import List
try:
    from .schemas import Item, ItemCreate
except ImportError:
    from schemas import Item, ItemCreate

items: List[Item] = []
next_id = 1

def get_items():
    return items

def create_item(item: ItemCreate):
    global next_id
    new_item = Item(id=next_id, value=item.value)
    items.append(new_item)
    next_id += 1
    return new_item

def update_item(item_id: int, item: ItemCreate):
    for i, existing_item in enumerate(items):
        if existing_item.id == item_id:
            updated_item = Item(id=item_id, value=item.value)
            items[i] = updated_item
            return updated_item
    return None

def delete_item(item_id: int):
    for i, existing_item in enumerate(items):
        if existing_item.id == item_id:
            del items[i]
            return True
    return False 