import React, { useEffect, useState } from 'react';

const Reservas = () => {
  const [bookings, setBookings] = useState([]);

  useEffect(() => {
    // Verificar si el usuario es un administrador al cargar la página
    const isAdmin = localStorage.getItem("isAdmin");
    if (isAdmin === "true") {
      // Si el usuario es un administrador, obtener las reservas de todos los usuarios
      fetchBookings();
    } else {
      // Si el usuario no es un administrador, redirigir a otra página o mostrar un mensaje de acceso denegado
      // Ejemplo: navigate("/acceso-denegado");
      console.log("Acceso denegado");
    }
  }, []);

  const fetchBookings = async () => {
    try {
      // Realizar la solicitud para obtener las reservas de todos los usuarios
      const response = await fetch('http://localhost:8090/booking');
      const data = await response.json();

      // Almacenar las reservas en el estado
      setBookings(data.bookings);
    } catch (error) {
      console.log('Error al obtener las reservas:', error);
    }
  };

  return (
    <div>
      <h2>Reservas de todos los usuarios</h2>
      <ul>
        {bookings.map((booking) => (
          <li key={booking.id}>
            {/* Mostrar los detalles de la reserva */}
            <p>ID: {booking.id}</p>
            <p>Hotel ID: {booking.hotel_id}</p>
            <p>User ID: {booking.user_id}</p>
            <p>Fecha de inicio: {booking.startdate}</p>
            <p>Fecha de fin: {booking.enddate}</p>
            <p>Precio total: {booking.total_price}</p>
            {/* Agregar más detalles si es necesario */}
          </li>
        ))}
      </ul>
    </div>
  );
};

export default Reservas;