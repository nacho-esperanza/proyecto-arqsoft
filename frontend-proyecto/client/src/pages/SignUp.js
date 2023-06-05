import React from "react";

import './SignUp.css';

const SignUp = () => {
 
  return (

    


<div className="signup-container-1">

<h1>Caminos del Viento</h1>

<div className="signup-container">

<h2>
        Registro de Usuario
</h2>
     

<form>

<input type="text" placeholder="Nombre"/>

<input type="texto" placeholder="Apellido"/>

<input type="email" placeholder="Email"/>

<input type="password" placeholder="Contraseña"/>

<input type="password" placeholder="Repita Contraseña"/>

<button type="submit">Registrate</button>

</form>

<p>Ya tengo Usuario</p>

<button type="submit" onClick={() => window.location.href = '/login'}>Iniciar Sesion</button>
<button onClick={() => window.location.href = '/#'}>Back</button>

    </div>
</div>
  );
};

export default SignUp;
