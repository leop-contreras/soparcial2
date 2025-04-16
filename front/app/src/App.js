import React, { useState, useEffect } from 'react';
import './App.css'; // Add styles here if needed

function App() {
  const [users, setUsers] = useState([]);

  const [id, setId] = useState('');
  const [name, setName] = useState('');
  const [email, setEmail] = useState('');

  useEffect(() => {
    // Agarra la info de la API
    fetch('http://localhost:8000/getusers',{
      method: "GET",
      })
      .then((response) => response.json())
      .then((data) => {setUsers(data);})
      .catch((error) => {console.error('Error fetching data:', error);});
  }, []);

  const AddUser = () => {
    if(isNaN(id) || id <= 0){
      console.warn("Invalid ID");
      setId('');
      alert('No es un ID válido');
      return;
    }

    for (let user of users) {
      if (Number(user.id) === Number(id)) {
        console.warn("Duplicate ID");
        setId('');
        alert('ID duplicado');
        return;
      }
    }

    fetch('http://localhost:8000/adduser', {
      method: "POST",
      headers: {"Content-Type": "application/json"},
      body: JSON.stringify({
        id: Number(id),
        name: name,
        email: email
      }),
    })
      .then((response) => response.json())
      .then((data) => {
        console.log(data);

        fetch('http://localhost:8000/getusers',{
          method: "GET",
        })
        .then((response) => response.json())
        .then((data) => {setUsers(data);})
        .catch((error) => {console.error('Error fetching data:', error);});

        setId('');
        setName('');
        setEmail('');
      
      })
      .catch((error) => {console.error(error);});
  };

  const RemoveUser = () => {
    if(isNaN(id) || id <= 0){
      console.warn("Invalid ID");
      setId('');
      alert('No es un ID válido');
      return;
    }

    fetch('http://localhost:8000/removeuser', {
      method: "POST",
      headers: {"Content-Type": "application/json"},
      body: JSON.stringify({
        id: Number(id)
      }),
    })
      .then((response) => response.json())
      .then((data) => {
        console.log(data);

        fetch('http://localhost:8000/getusers',{
          method: "GET",
        })
        .then((response) => response.json())
        .then((data) => {setUsers(data);})
        .catch((error) => {console.error('Error fetching data:', error);});

        setId('');
        setName('');
        setEmail('');
      
      })
      .catch((error) => {console.error(error);});
    
    fetch('http://localhost:8000/getusers',{
        method: "GET",
    })
    .then((response) => response.json())
    .then((data) => {setUsers(data);})
    .catch((error) => {console.error('Error fetching data:', error);});
  };

  return (
    <div className="appContainer">
      <h1 className="title">Usuarios</h1>
      <ul className="userList">
        {users.map((user) => (
          <li key={user.id} className="userItem">
            <span className="userID">#{user.id}</span> — <span className="user-detail">{user.name}</span> — <span className="user-detail">{user.email}</span>
          </li>
        ))}
      </ul>

      <div className='crudContainer'>
        <div>
          <input id="idInput" placeholder='ID' value={id} onChange={(e) => setId(e.target.value)}></input>
          <input id="nameInput" placeholder='Name' value={name} onChange={(e) => setName(e.target.value)}></input>
          <input id="emailInput" placeholder='Email' value={email} onChange={(e) => setEmail(e.target.value)}></input>
        </div>
        <div id="crudButtonContainer">
          <button id='addUser' onClick={AddUser}>Añadir</button>
          <button id='removeUser' onClick={RemoveUser}>Eliminar (solo ID)</button>
        </div>
      </div>
    </div>
  );
}

export default App;
