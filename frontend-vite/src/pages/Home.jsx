import React, { useEffect, useState } from 'react';
import { useParams } from 'react-router-dom';
import {useNavigate} from "react-router-dom";


import './Home.css';
import '../App.css';
import Barra from "../componentes/Barra/Barra";


const Home = () => {
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
    return <div>Cargando hoteles por favor espere...</div>;
  }



  return (
  <div className="Planas">
  <Barra />

  <div class="hoteles">

  <div class="titulo">
    <h2>{hotel.name}</h2>
  </div>
  <div class="descripcion">
     <p>{hotel.description}</p>
  </div>

  <div class="imagen">
  <img src={image.url} alt={`Imagen ${index}`} />
  </div>
  </div>
</div>
      
    );
  };  



  

    
  export default Home;