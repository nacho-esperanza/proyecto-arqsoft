import "./Barra.css"

const Barra = () => {
  return (
    <div className="NaveBa">
      <div className="NaveC">
        <span className="logo">SuperCool</span>
        <div className="NaveIt">
          <button className="NaveBo">Register</button>
          <button className="NaveBo" href="/login">Login</button>
          <a href="/login">este anda?</a>
        </div>
      </div>
    </div>
  )
}

export default Barra