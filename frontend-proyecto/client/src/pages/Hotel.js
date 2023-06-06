import React from "react";
import axios from "axios";

import './Hotel.css';  

const Hotel = () => {

  const GetHotelByid = async (id) => {
    try {
      const response = await axios.get(`/user/:id`);
      console.log(response.data); 
    } catch (error) {
      console.error(error);
    }
  };

  return (
    <div className="home-container">
    
    
    <div className="menu_titulo">
    <p>$id{ }</p>
    <h1>Caminos del Viento</h1>

    </div>



    <div className="nombre_hotel">Nombre del Hotel</div>
    <div className="hotel_descripcion">Descripción del Hotel
    
    <p>Aquí va la descripción del hotel. Puedes agregar información sobre las características, servicios y ubicación del hotel.Proporciona detalles atractivos y convincentes para que los usuarios se sientan interesados en hospedarse en tu hotel.Aquí va la descripción del hotel. Puedes agregar información sobre las características, servicios y ubicación del hotel.Proporciona detalles atractivos y convincentes para que los usuarios se sientan interesados en hospedarse en tu hotel.Aquí va la descripción del hotel. Puedes agregar información sobre las características, servicios y ubicación del hotel.Proporciona detalles atractivos y convincentes para que los usuarios se sientan interesados en hospedarse en tu hotel.Aquí va la descripción del hotel. Puedes agregar información sobre las características, servicios y ubicación del hotel.Proporciona detalles atractivos y convincentes para que los usuarios se sientan interesados en hospedarse en tu hotel.   </p>
    
    </div>
    
    <div className="hotel_img">Imagen del Hotel
    
    </div>
    
    
    <div className="hotel_caracteristica">Características del Hotel:
    <li>Caracteristica</li>
    <li>Caracteristica</li>
    <li>Caracteristica</li>
    <li>Caracteristica</li>
    </div>
    <div className="hotel_Precio">Precio del Hotel
    <br></br>
    <br></br>
    <button>Reservar</button>
    </div>
  </div>
  );
};
  
export default Hotel;