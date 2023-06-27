import { useState } from 'react';
import axios from 'axios';

const CrearHotel = () => {
  const [hotel, setHotel] = useState({
    name: '',
    address: '',
    city: '',
    stars: 0,
    description: '',
    price: 0,
    rooms: 0,
    parking: false,
    pool: false,
    wifi: false,
    air: false,
    gym: false,
    spa: false,
  });

  const handleInputChange = (event) => {
    const { name, value } = event.target;
    setHotel((prevHotel) => ({
      ...prevHotel,
      [name]: value,
    }));
  };

  const handleCheckboxChange = (event) => {
    const { name, checked } = event.target;
    setHotel((prevHotel) => ({
      ...prevHotel,
      [name]: checked,
    }));
  };

  const handleSubmit = async (event) => {
    event.preventDefault();

    try {
      const response = await axios.post('http://localhost:8090/hotel', hotel);
      // Manejar la respuesta exitosa
      console.log('Hotel creado:', response.data);
    } catch (error) {
      // Manejar el error
      console.error('Error al crear el hotel:', error);
    }
  };

  return (
    <div>
      <h1>Crear Hotel</h1>
      <form onSubmit={handleSubmit}>
        <div>
          <label>Nombre:</label>
          <input type="text" name="name" value={hotel.name} onChange={handleInputChange} />
        </div>
        <div>
          <label>Dirección:</label>
          <input type="text" name="address" value={hotel.address} onChange={handleInputChange} />
        </div>
        <div>
          <label>Ciudad:</label>
          <input type="text" name="city" value={hotel.city} onChange={handleInputChange} />
        </div>
        <div>
          <label>Estrellas:</label>
          <input type="number" name="stars" value={hotel.stars} onChange={handleInputChange} />
        </div>
        <div>
          <label>Descripción:</label>
          <textarea name="description" value={hotel.description} onChange={handleInputChange} />
        </div>
        <div>
          <label>Precio:</label>
          <input type="number" name="price" value={hotel.price} onChange={handleInputChange} />
        </div>
        <div>
          <label>Habitaciones:</label>
          <input type="number" name="rooms" value={hotel.rooms} onChange={handleInputChange} />
        </div>
        <div>
          <label>Parking:</label>
          <input
            type="checkbox"
            name="parking"
            checked={hotel.parking}
            onChange={handleCheckboxChange}
          />
        </div>
        <div>
          <label>Piscina:</label>
          <input type="checkbox" name="pool" checked={hotel.pool} onChange={handleCheckboxChange} />
        </div>
        <div>
          <label>Wifi:</label>
          <input type="checkbox" name="wifi" checked={hotel.wifi} onChange={handleCheckboxChange} />
        </div>
        <div>
          <label>Aire acondicionado:</label>
          <input type="checkbox" name="air" checked={hotel.air} onChange={handleCheckboxChange} />
        </div>
        <div>
          <label>Gimnasio:</label>
          <input type="checkbox" name="gym" checked={hotel.gym} onChange={handleCheckboxChange} />
        </div>
        <div>
          <label>Spa:</label>
          <input type="checkbox" name="spa" checked={hotel.spa} onChange={handleCheckboxChange} />
        </div>
        <button type="submit">Crear Hotel</button>
      </form>
    </div>
  );
};

export default CrearHotel;
