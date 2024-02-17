import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import './SignUp.css';
import toast, { Toaster } from 'react-hot-toast';

const SignUp = () => {
  const navigate = useNavigate();
  const [user, setUser] = useState({
    name: '',
    lastname: '',
    email: '',
    password: '',
  });

  const handleChange = (e) => {
    const { name, value } = e.target;
    setUser(prevState => ({
      ...prevState,
      [name]: value
    }));
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    const { name, lastname, email, password } = user;

    // Validación simplificada
    if (!name || !lastname || !email || !password) {
      toast.error('Todos los campos son obligatorios');
      return;
    }

    try {
      const response = await fetch('http://localhost:8090/user', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(user),
      });

      if (response.ok) {
        toast.success('Usuario creado');
        setTimeout(() => navigate('/'), 2000);
      } else {
        toast.error('Error al crear usuario');
      }
    } catch (error) {
      console.error('Error al realizar la solicitud:', error);
    }
  };

  return (
    <div className="signup-container-1">
      <div className="signup-container">
        <h2>Registro de Usuario</h2>
        <form id="formRegister" onSubmit={handleSubmit}>
          <input
            name="name"
            type="text"
            placeholder="Nombre"
            value={user.name}
            onChange={handleChange}
          />
          <input
            name="lastname"
            type="text"
            placeholder="Apellido"
            value={user.lastname}
            onChange={handleChange}
          />
          <input
            name="email"
            type="email"
            placeholder="Email"
            value={user.email}
            onChange={handleChange}
          />
          <input
            name="password"
            type="password"
            placeholder="Contraseña"
            value={user.password}
            onChange={handleChange}
          />
          <input type="password" placeholder="Repita Contraseña" />
          <div className='botonRegistrarse'>
          <button id="botonRegistrarse"  type="submit">Registrate</button>
          </div>
        </form>
        <Toaster position="top-center" reverseOrder={false} />

        <div className="signup-actions">
          <p>Ya tengo Usuario, <span className="link" onClick={() => navigate("/login")}>Iniciar Sesión</span></p>
        </div>
      </div>
    </div>
  );
};

export default SignUp;
