import React, { useEffect, useState } from 'react';

import './Admin.css';

const Admin = () => {
  const isAdmin = localStorage.getItem("isAdmin");

  useEffect(() => {
    // Verificar si el usuario es un administrador
    if (isAdmin !== "true") {
      // Si no es un administrador, redirigir a otra página
      window.location.href = '/';
    }
  }, [isAdmin]);

  const handleReservas = () => {
    window.location.href = '/reservas';
  };

  const handleCrearHotel = () => {
    window.location.href = '/crearhotel';
  };

  return (
    <div className="container">
      <h1>Página de Administrador</h1>
      <button onClick={handleReservas}>Ver Reservas</button>
      <button onClick={handleCrearHotel}>Crear Hotel</button>
    </div>
  );
};
export default Admin;