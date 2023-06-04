import "./Barra.css"

const Barra = () => {
  return (
    <div className="NaveBa">
      <div className="NaveC">
        <h2 className="logo">Hoteles kicillof</h2>
        <div className="NaveIt">
          <button className="NaveBo">Registrar</button>
          <button className="NaveBo" onClick={() => window.location.href = '/login'}>Iniciar sesion</button>
        </div>
      </div>
    </div>
  )
}

export default Barra