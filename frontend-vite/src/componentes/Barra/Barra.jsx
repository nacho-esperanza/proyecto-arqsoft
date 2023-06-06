import "./Barra.css"

const Barra = () => {
  return (
    <div className="NaveBa">
      <div className="NaveC">
        <h1 className="logo">Caminos del Viento</h1>
        <div className="NaveIt">
          <button className="NaveBo"  onClick={() => window.location.href = '/signup'}>Registrar</button>
          <button className="NaveBo" onClick={() => window.location.href = '/login'}>Iniciar sesion</button>
        </div>
      </div>
    </div>
  )
}

export default Barra