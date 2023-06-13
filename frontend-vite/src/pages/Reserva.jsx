import React, {useEffect, useState} from 'react';
import DateRangeComp from '../componentes/DateRangeComp';
import {useNavigate, useParams} from 'react-router-dom';

import './Reserva.css';

// Import react hot toast
import toast, {Toaster} from 'react-hot-toast';

const Reserva = () => {
    const { id } = useParams();
    const [hotel, setHotel] = useState(null);
    const [startDate, setStartDate] = useState(null);
    const [endDate, setEndDate] = useState(null);
    const [precioTotal, setPrecioTotal] = useState("");
    const [errorMessage, setErrorMessage] = useState('');
    const navigate = useNavigate();

    const handleSubmit = async (e) => {
        e.preventDefault();

        let inputValue = e.target[0].value;
        const splitValues = inputValue.split(' - ');
        if (splitValues.length > 1) {
          try {
            const startDateString = splitValues[0];
            const endDateString = splitValues[1];
            let tempStartDate = startDateString.replace(/\//g, '-');
            let tempEndDate = endDateString.replace(/\//g, '-');
            tempStartDate += "T00:00:00-03:00";
            tempEndDate += "T00:00:00-03:00";
            setStartDate(tempStartDate);
            setEndDate(tempEndDate);


            const response = await fetch('http://localhost:8090/booking', {
              method: 'POST',
              headers: {
                'Content-Type': 'application/json',
              },
              body: JSON.stringify({
                hotel_id: hotel.id,
                user_id: Number(localStorage.getItem("user_id")),
                startdate: tempStartDate,
                enddate: tempEndDate
              })
            });
            if (response.ok) {
                const data = await response.json();
        setPrecioTotal(data.totalPrice);

              toast.success('Se ha registrado su reserva correctamente!');      

              console.log('Se ha registrado su reserva');
            } else {
              toast.error('No hay habitaciones disponibles');
              console.log('No hay habitaciones disponibles');
              
              // Mostrar por consola el json enviado al backend
            console.log(JSON.stringify({ 
                hotel_id: hotel.id, 
                user_id: Number(localStorage.getItem("user_id")), 
                startdate: tempStartDate, 
                enddate: tempEndDate 
            }));
            }
          } catch (error) {
            console.log('Error al realizar la solicitud al backend:', error);
          }
        }
      };

      

      useEffect(() => {
        const fetchHotel = async () => {
          try {
            const response = await fetch(`http://localhost:8090/hotel/id/${id}`);
            const data = await response.json();
            setHotel(data);
            console.log(data)
          } catch (error) {
            console.log('Error al obtener el hotel:', error);
          }
        };
    
        fetchHotel();
      }, [id]);

      useEffect(() => {
        // Calcula el precio total cuando las fechas de check-in y check-out se actualizan
        const calculateTotalPrice = () => {
          if (startDate && endDate && hotel) {
            // Calcula la diferencia en días entre las fechas de check-in y check-out
            const diffInTime = new Date(endDate) - new Date(startDate);
            const diffInDays = Math.ceil(diffInTime / (1000 * 60 * 60 * 24));
    
            // Calcula el precio total multiplicando el precio diario del hotel por la cantidad de días
            const total = hotel.price * diffInDays;
    
            // Actualiza el estado del precio total
            console.log(total);
            setPrecioTotal(total);
          }
        };
    
        calculateTotalPrice();
      }, [startDate, endDate, hotel]);

      return (
        <div>
      {hotel ? (
        <div className="home-container">
          <div >
            <div className="menu_titulo">
                <h1>Reserva</h1>
            </div>
            <div className="nombre_hotel">{hotel.name}</div>
                <div className="hotel_descripcion">
                    Descripción del Hotel
                    <p>{hotel.description}</p>
                </div>

                <div className="hotel_fechas">
                
                    <h3>Seleccione check-in y check-out:</h3>
                    <form id="formLogin" onSubmit={handleSubmit}>
                        <DateRangeComp onChange={(item) => { setStartDate(item.startDate); setEndDate(item.endDate); }} />
                        Precio Total: {precioTotal}
                        {errorMessage && <p style={{ color: 'red' }}>{errorMessage}</p>}
                        <button id="botonLogin" type="submit" >Reservar</button>
                    </form>
                    <Toaster/>
            
                 </div>

                 <div className="hotel_caracteristica">
        Características del Hotel:
        <ul>
          <li>Parking: {hotel.parking ? 'Incluye' : 'No Incluye'}</li>
          <li>Pool: {hotel.pool ? 'Incluye' : 'No Incluye'}</li>
          <li>Wifi: {hotel.wifi ? 'Incluye' : 'No Incluye'}</li>
          <li>Aire: {hotel.air ? 'Incluye' : 'No Incluye'}</li>
          <li>Gym: {hotel.gym ? 'Incluye' : 'No Incluye'}</li>
          <li>Spa: {hotel.spa ? 'Incluye' : 'No Incluye'}</li>
        </ul>
      </div>
      {/*<div className="hotel_Precio">
        Precio Total: {precioTotal}
        <br></br>
        <br></br>
        <button>Reservar</button>
      </div>*/}
      <div className='hotel_Precio'>
      <button  onClick={() => window.location.href = `/hotel/${hotel.id}`}>Volver</button>
      </div>

           
          </div>
          </div>
      ) : (
          <p>No se encontró el hotel</p>
      )}
    
    </div>
);
};

      
  
  export default Reserva;
