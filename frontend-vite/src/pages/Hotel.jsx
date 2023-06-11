/*
import React, { useEffect, useState } from 'react';
import { useParams } from 'react-router-dom';

import './Hotel.css';  

const Hotel = () => {
    const { id } = useParams();
    const [hotel, setHotel] = useState(null);

    // hotelsdata.map (repite todo el array de hoteles)
  
  
    useEffect(() => {
      const fetchHotel = async () => {
        try {
          const response = await fetch(`http://localhost:8090/hotel/id/${id}`);
          const data = await response.json();
          setHotel(data);
        } catch (error) {
          console.log('Error al obtener el hotel:', error);
        }
      };
  
      fetchHotel();
    }, [id]);



    return (
        <div className="home-container">
        
        
        <div className="menu_titulo">
        <p>Username</p>
        <h1>Caminos del Viento</h1>
    
        </div>
    
    
    
        <div className="nombre_hotel">{hotel.name}</div>
        <div className="hotel_descripcion">Descripción del Hotel
        
        <p> {hotel.description}  </p>
         
        </div>
        
        <div className="hotel_img">Imagen del Hotel
        
        </div>
        
        
        <div className="hotel_caracteristica">Características del Hotel:
        <li>Parking: {hotel.parking}</li>
        <li>Pool: {hotel.pool}</li>
        <li>Wifi: {hotel.wifi}</li>
        <li>Aire: {hotel.air}</li>
        <li>Gym: {hotel.gym}</li>
        <li>Spa: {hotel.spa}</li>
        </div>
        <div className="hotel_Precio">Precio del Hotel: {hotel.price}
        <br></br>
        <br></br>
        <button>Reservar</button>
        </div>
      </div>
      );
    };
      
    export default Hotel;*/


// ESTO ES EL MISMO CODIGO QUE ARRIBA PERO CORREGIDO

import React, { useEffect, useState } from 'react';
import { useParams } from 'react-router-dom';
import './Hotel.css';

const Hotel = () => {
  const { id } = useParams();
  const [hotel, setHotel] = useState(null);
  const [images, setImages] = useState([]);

  // hotelsdata.map (repite todo el array de hoteles)

  useEffect(() => {
    const fetchHotel = async () => {
      try {
        const response = await fetch(`http://localhost:8090/hotel/id/${id}`);
        const data = await response.json();
        setHotel(data);
      } catch (error) {
        console.log('Error al obtener el hotel:', error);
      }
    };

    const fetchImages = async () => {
      try {
        const response = await fetch(`http://localhost:8090/hotel/${id}/images/`);
        const data = await response.json();
        setImages(data);
        console.log(images)
      } catch (error) {
        console.log('Error al obtener las imágenes del hotel:', error);
      }
    };

    fetchImages();
    fetchHotel();
  }, [id]);
  

  if (!hotel) {
    return <div>Cargando hotel...</div>;
  }

  return (
    <div className="home-container">
      <div className="menu_titulo">
        <p>Username</p>
        <h1>Caminos del Viento</h1>
      </div>

      <div className="nombre_hotel">{hotel.name}</div>
      <div className="hotel_descripcion">
        Descripción del Hotel
        <p>{hotel.description}</p>
      </div>


      // ARREGLEN LAS IMAGENES POR FAVOR
      <div className="hotel_img">
        {images && images.map((image, index) => (
          <img src={image.url} alt={`Imagen ${index}`} key={index} />
        ))}
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
      <div className="hotel_Precio">
        Precio del Hotel: {hotel.price}
        <br></br>
        <br></br>
        <button>Reservar</button>
      </div>
    </div>
  );
};

export default Hotel;