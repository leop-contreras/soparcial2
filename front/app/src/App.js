import './App.css';
import React, { useState, useEffect } from 'react';

function App() {
  const [users, setUsers] = useState([]);

  useEffect(() => {
    // Agarra la info de la API
    fetch('http://localhost:8000/data')
      .then((response) => response.json())
      .then((data) => {
        setUsers(data);
      })
      .catch((error) => {
        console.error('Error fetching data:', error);
      });
  }, []);

  return (
    <div>
      <h1>users</h1>
      <ul>
        {users.map((user) => (
          <li key={user.id}>
            {user.id}: {user.name}
          </li>
        ))}
      </ul>
    </div>
  );
}

export default App;
