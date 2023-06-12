import React, { useEffect, useState } from 'react';
import { useParams } from 'react-router-dom';
import './MisReservas.css';

const MisReservas = () => {
    // Obtener el id del usuario del localStorage
    const id = localStorage.getItem('user_id');

    const [bookings, setBookings] = useState([]);
  
    useEffect(() => {
      const fetchBookings = async () => {
        try {
          // Hacer la solicitud al backend para obtener las reservas del cliente
          const response = await fetch(`http://localhost:8090/booking/user/${id}`, {
            headers: {
              Authorization: `Bearer ${localStorage.getItem('token')}`, // Agrega el token de autenticación en los headers
            },
           
          });
  
          if (response.ok) {
            console.log(response)
            const data = await response.json();
             // console log de la respuesta
             console.log(data)
             
             // console log del largo de la respuesta
                console.log(data.bookings.length)
                console.log(bookings); // Verifica el contenido del array bookings
            setBookings(data.bookings);
          } else {
            console.log('Error al obtener las reservas');
          }
        } catch (error) {
          console.log('Error al realizar la solicitud al backend:', error);
        }
      };
  
      fetchBookings();
    }, []);
  
    return (
      <div className='reservas'>
        <h1>Mis Reservas</h1>
        {bookings.length > 0 ? (
          <table>
            <thead>
              <tr>
                <th>ID de Reserva</th>
                <th>ID de Usuario</th>
                <th>Hotel</th>
                <th>Fecha de Entrada</th>
                <th>Fecha de Salida</th>
                <th>Precio Final</th>
              </tr>
            </thead>
            <tbody>
              {bookings.map((booking) => (
                <tr key={booking.id}>
                  <td>{booking.id}</td>
                    <td>{booking.user_id}</td>
                  <td>{booking.hotel_id}</td>
                  <td>{booking.startdate}</td>
                  <td>{booking.enddate}</td>
                  <td>{booking.total_price}</td>
                </tr>
              ))}
            </tbody>
          </table>
        ) : (
          <p>No se encontraron reservas.</p>
        )}
      </div>
    );
  };
  
  export default MisReservas;