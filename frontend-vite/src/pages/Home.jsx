import React, { useEffect, useState } from 'react';
import { useParams } from 'react-router-dom';
import { useNavigate } from 'react-router-dom';
import { Carousel } from 'react-responsive-carousel';

import './Home.css';
import '../App.css';
import Barra from '../componentes/Barra/Barra';
import 'react-responsive-carousel/lib/styles/carousel.min.css';

const Home = () => {
  const { id } = useParams();
  const [hotels, setHotels] = useState([]);
  const [images, setImages] = useState([]);

  useEffect(() => {
    const fetchHotels = async () => {
      try {
        const response = await fetch('http://localhost:8090/hotel');
        const data = await response.json();
        setHotels(data);
        console.log(data);
      } catch (error) {
        console.log('Error al obtener los hoteles:', error);
      }
    };

    const fetchImages = async () => {
      try {
        const response = await fetch('http://localhost:8090/hotel/images');
        const data = await response.json();
        setImages(data);
        console.log(data);
      } catch (error) {
        console.log('Error al obtener las imágenes de los hoteles:', error);
      }
    };

    fetchHotels();
    fetchImages();
  }, []);

  if (!hotels.length || !images.length) {
    return <div>Cargando hoteles, por favor espere...</div>;
  }

  return (
    <div className="Planas">
      <Barra />

      <div className="hoteles">
        {hotels.map((hotel, index) => (
          <div key={index}>
            <div className="titulo">
              <h2>{hotel.name}</h2>
            </div>
            <div className="descripcion">
              <p>{hotel.description}</p>
            </div>
            <div className="imagen">
              <Carousel showThumbs={false}>
                {images
                  .filter((image) => image.hotel_id === hotel.id)
                  .map((image, i) => (
                    <div key={i}>
                      <img src={image.url} alt={`Imagen ${i}`} />
                    </div>
                  ))}
              </Carousel>
            </div>
            <div className='precio'>
              <h3>Precio por noche: <h2>US$ {hotel.price}</h2></h3>
            </div>
            <div className="boton">
            <button  onClick={() => window.location.href = `/hotel/${hotel.id}`}>Ver Hotel</button>
            </div>
          </div>
        ))}
      </div>
    </div>
  );
};

export default Home;

