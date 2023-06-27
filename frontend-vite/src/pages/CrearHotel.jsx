import { useState } from 'react';
// importar hot toaster
import toast, { Toaster } from 'react-hot-toast';


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

    if (hotel.name === '') {
        document.getElementById('inputNombreHotel').style.borderColor = 'red';
      } else if (hotel.address === '') {
        document.getElementById('inputDireccionHotel').style.borderColor = 'red';
      } else if (hotel.city === '') {
        document.getElementById('inputCiudadHotel').style.borderColor = 'red';
      } else if (hotel.stars === 0) {
        document.getElementById('inputEstrellasHotel').style.borderColor = 'red';
      } else if (hotel.description === '') {
        document.getElementById('inputDescripcionHotel').style.borderColor = 'red';
      } else if (hotel.price === 0) {
        document.getElementById('inputPrecioHotel').style.borderColor = 'red';
      } else if (hotel.rooms === 0) {
        document.getElementById('inputHabitacionesHotel').style.borderColor = 'red';
      } else {
    try {
        const response = await fetch( 'http://localhost:8090/hotel', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(hotel),
        }
            )
        if (response.ok) {
            console.log('Hotel creado');
            toast.success('Hotel creado');
            setTimeout(() => {
                navigate('/home');
            }, 2000); // Redirige después de 2 segundos a la página de inicio
        } else {
            toast.error('Hotel inválido');
            console.log('Hotel inválido');
        }
    } catch (error) {
        console.log(error);
    }
    }


    // Enviar datos al backend



    
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
      <Toaster position="top-center" reverseOrder={false} />
    </div>
  );
};

export default CrearHotel;
