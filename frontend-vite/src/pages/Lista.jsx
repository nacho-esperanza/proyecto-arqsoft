import React from "react";

import '../App.css';
  
const Lista = () => {
    const [hoteles, setHoteles] = useState([]);
    const navigate = useNavigate();
    const selectHotels = (id) => {
        navigate(`/hotel/${id}`);
    };

    useEffect(() => {
        // Realizar la solicitud al backend para obtener la lista de hoteles
        const fetchHoteles = async () => {
            try {
                const response = await fetch('http://localhost:8080/hotel');
                const data = await response.json();
                setHoteles(data);
            } catch (error) {
                console.log('Error al obtener la lista de hoteles:', error);
            }
        };

        fetchHoteles();
    }, []);

  return (
   
    <div style={{ alignItems: 'left', backgroundColor: '#CBE4DE', minHeight: '100vh' }}>
    <h1 style={{ textAlign: 'center', color:'#0E8388'}}>Hoteles:</h1>
    {hoteles.length > 0 ? (
        <div>
            {hoteles.map((hotel) => (
                <div key={hotel.id} style={{ display: 'flex', alignItems: 'center', justifyContent: 'flex-start', marginBottom: '20px' }}>
                    <img src={hotel.imagen} style={{ width: '150px', height: '150px', marginRight: '10px', marginLeft:'30px' }} />
                    <div>
                        <h2 style={{color: '#0E8388' }}>{hotel.name}</h2>
                        <p style={{ color: '#2C3333' }}>Estrellas: {hotel.stars}</p>
                        <p style={{ color: '#2C3333', marginRight: 'auto' }}>Precio por noche: ${hotel.price}</p>
                    </div>
                    <button style={{ marginRight: 'auto', backgroundColor:'#2E4F4F' }} type="submit" onClick={() => selectHotels(hotel.id)}>Ver</button>

                </div>
            ))}
        </div>
    ) : (
        <p>No se encontraron hoteles.</p>
    )}
</div>

);
};
  
export default Lista;