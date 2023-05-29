import "./Barra.css"

const Barra = () => {
  return (
    <div className="NaveBa">
      <div className="NaveC">
        <span className="logo">SuperCool</span>
        <div className="NaveIt">
          <button className="NaveBo">Register</button>
          <button className="NaveBo" onClick={() => window.location.href = '/login'}>Login</button>
        </div>
      </div>
    </div>
  )
}

export default Barra