
import React, { useEffect, useState } from 'react';

import "./Barra.css";

const Barra = () => {
  const userId = localStorage.getItem("user_id");
  const token = localStorage.getItem("token");
  const [user, setUser] = useState(null);
  const isAdmin = localStorage.getItem("isAdmin");

  useEffect(() => {
    const fetchData = async () => {
      const response = await fetch(`http://localhost:8090/user/${userId}`);
      if (response.ok) {
        const data = await response.json();
        setUser(data);
      }
    };

    fetchData();
  }, [userId]);
  const handleLogin = () => {
    window.location.href = '/login';
  };

  const handleSignup = () => {
    window.location.href = '/signup';
  };

  const handleLogout = () => {
    localStorage.removeItem("token");
    localStorage.removeItem("user_id");
    window.location.href = '/';
  };

  return (
    <div className="NaveBa">
      <div className="NaveC">
        <h1>SaNaFrAu</h1>
        
        <div className="NaveIt">
        {user && (
            <>
             
              {isAdmin === "true" && (
                <>
                  {/* Mostrar opciones específicas para el administrador */}
                  <button className="NaveBo" onClick={() => window.location.href = '/admin'}>Panel de Administrador</button>
                  <button className="NaveBo" name='button_SingUp' onClick={() => window.location.href = '/'}>Home</button>
                </>
              )}
              {isAdmin === "false" && (
                <>
                  {/* Mostrar opciones específicas para el cliente */}
                  <button className="NaveBo" onClick={() => window.location.href = '/misreservas'}>Mis Reservas</button>
                  <button className="NaveBo" name='button_SingUp' onClick={() => window.location.href = '/'}>Home</button>
                </>
              )}
              <button className="NaveBo" onClick={handleLogout}>Cerrar sesión</button>
            </>
          )}
          {!user && (
            <>
              {/* Mostrar opciones para un usuario no autenticado */}
              <button className="NaveBo" onClick={handleSignup}>Registrar</button>
              <button className="NaveBo" onClick={handleLogin}>Iniciar</button>
              <button className="NaveBo" name='button_SingUp' onClick={() => window.location.href = '/'}>Home</button>
            </>
          )}

        </div>
      </div>
    </div>
  );
};

export default Barra;