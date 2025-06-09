import React, { useState, useEffect } from 'react';
import './App.css';

const API_URL = 'http://localhost:8000/items';

function App() {
  const [items, setItems] = useState([]);
  const [input, setInput] = useState('');
  const [editIndex, setEditIndex] = useState(null);
  const [editValue, setEditValue] = useState('');

  useEffect(() => {
    fetch(API_URL)
      .then(res => res.json())
      .then(data => setItems(data));
  }, []);

  const handleAdd = async () => {
    if (input.trim() !== '') {
      const res = await fetch(API_URL, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ value: input })
      });
      const newItem = await res.json();
      setItems([...items, newItem]);
      setInput('');
    }
  };

  const handleDelete = async (id) => {
    await fetch(`${API_URL}/${id}`, { method: 'DELETE' });
    setItems(items.filter(item => item.id !== id));
  };

  const handleEdit = (index) => {
    setEditIndex(index);
    setEditValue(items[index].value);
  };

  const handleUpdate = async () => {
    if (editValue.trim() !== '') {
      const id = items[editIndex].id;
      const res = await fetch(`${API_URL}/${id}`, {
        method: 'PUT',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ value: editValue })
      });
      const updatedItem = await res.json();
      const updatedItems = [...items];
      updatedItems[editIndex] = updatedItem;
      setItems(updatedItems);
      setEditIndex(null);
      setEditValue('');
    }
  };

  return (
    <div className="App">
      <h1>Simple CRUD App (with FastAPI backend)</h1>
      <div style={{ marginBottom: 20 }}>
        <input
          type="text"
          value={input}
          onChange={e => setInput(e.target.value)}
          placeholder="Add new item"
        />
        <button onClick={handleAdd}>Add</button>
      </div>
      <ul>
        {items.map((item, index) => (
          <li key={item.id} style={{ marginBottom: 10 }}>
            {editIndex === index ? (
              <>
                <input
                  type="text"
                  value={editValue}
                  onChange={e => setEditValue(e.target.value)}
                />
                <button onClick={handleUpdate}>Update</button>
                <button onClick={() => setEditIndex(null)}>Cancel</button>
              </>
            ) : (
              <>
                {item.value}
                <button onClick={() => handleEdit(index)} style={{ marginLeft: 10 }}>Edit</button>
                <button onClick={() => handleDelete(item.id)} style={{ marginLeft: 5 }}>Delete</button>
              </>
            )}
          </li>
        ))}
      </ul>
    </div>
  );
}

export default App;
