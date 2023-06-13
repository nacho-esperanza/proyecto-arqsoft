
import React, { useEffect, useState } from 'react';
import { useParams } from 'react-router-dom';
import './Hotel.css';

// librerias para el carrusel de imagenes
import Slider from 'react-slick';
import 'slick-carousel/slick/slick.css';
import 'slick-carousel/slick/slick-theme.css';

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

    // Para las imagenes es necesario activar/desactivar CORS en el navegador
    const fetchImages = async () => {
      try {
        const response = await fetch(`http://localhost:8090/hotel/${id}/images/`);
        const data = await response.json();
        setImages(data);
        console.log(data)
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
        <h1>Caminos del Viento</h1>
      </div>

      <div className="nombre_hotel">
        {hotel.name}
        </div>
      <div className="hotel_descripcion">
        Descripción del Hotel
        <p>{hotel.description}</p>
        <br />
        <p>Direccion del Hotel: {hotel.adress}</p>
      </div>


      <div className="hotel_img">
      <Slider dots={false} infinite={true} slidesToShow={1} slidesToScroll={1} className="slick-slider">
         {images &&
            images.map((image, index) => (
            <div className="slick-slide" key={index}>
              <img src={image.url} alt={`Imagen ${index}`} />
            </div>
        ))}
      </Slider>
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
        Precio del Hotel: <h2>US$ {hotel.price}</h2>
        <button  onClick={() => window.location.href = `/reserva/${hotel.id}`}>Reservar</button>
        <button onClick={() => window.location.href = `/`}> Volver</button>
      </div>
    </div>
  );
};

export default Hotel;