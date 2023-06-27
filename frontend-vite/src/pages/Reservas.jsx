import React, { useEffect, useState } from 'react';

const Reservas = () => {
  const [bookings, setBookings] = useState([]);
  const [hotels, setHotels] = useState([]);
  const [users, setUsers] = useState([]);

  useEffect(() => {
    // Verificar si el usuario es un administrador al cargar la página
    const isAdmin = localStorage.getItem("isAdmin");
    if (isAdmin === "true") {
      // Si el usuario es un administrador, obtener las reservas de todos los usuarios
      console.log ("El usuario es un administrador");
      fetchBookings();
      fetchHotels();
      fetchUsers();
    } else {
      // Si el usuario no es un administrador, redirigir a otra página o mostrar un mensaje de acceso denegado
      console.log("Acceso denegado");
    }
  }, []);

  const fetchBookings = async () => {
    try {
      // Realizar la solicitud para obtener las reservas de todos los usuarios
      const response = await fetch('http://localhost:8090/booking');
      const data = await response.json();

      // Almacenar las reservas en el estado
      console.log(data.bookings);
      setBookings(data.bookings);
    } catch (error) {
      console.log('Error al obtener las reservas:', error);
    }
  };

  const fetchHotels = async () => {
    try {
      // Realizar la solicitud para obtener los hoteles
      const response = await fetch('http://localhost:8090/hotel');
      const data = await response.json();

      // Almacenar los hoteles en el estado
      console.log(data.hotels);
      setHotels(data.hotels);
    } catch (error) {
      console.log('Error al obtener los hoteles:', error);
    }
  };

  const fetchUsers = async () => {
    try {
      // Realizar la solicitud para obtener los usuarios
      const response = await fetch('http://localhost:8090/user');
      const data = await response.json();
      
      // Almacenar los usuarios en el estado.
      console.log(data.users);
      setUsers(data.users);
    } catch (error) {
      console.log('Error al obtener los usuarios:', error);
    }
  };

  return (
    <div classname="hotel_caracteristica">
      <h2>Reservas de todos los usuarios</h2>
      <ul>
        {bookings.map((booking) => (
          <li key={booking.id}>
            {/* Mostrar los detalles de la reserva */}
            <p>ID: #{booking.id}</p>
            <p>ID del Hotel: {booking.hotel_id}</p>
            <p>ID del Usuario: {booking.user_id}</p>
            <p>Fecha de Ingreso: {booking.startdate}</p>
            <p>Fecha de Salida: {booking.enddate}</p>
            <p>Total Abonado: {booking.total_price}</p>
          </li>
        ))}
      </ul>
    </div>
  );
};

export default Reservas;