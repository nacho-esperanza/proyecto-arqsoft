import "./Barra.css"

const Barra = () => {
  return (
    <div className="NaveBa">
      <div className="NaveC">
        <h2 className="logo">SuperCool</h2>
        <div className="NaveIt">
          <button className="NaveBo">Register</button>
          <button className="NaveBo" onClick={() => window.location.href = '/login'}>Login</button>
        </div>
      </div>
    </div>
  )
}

export default Barra