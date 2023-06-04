import React from "react";

import './Home.css';
import '../App.css';
import Barra from "../componentes/Barra/Barra";
import Header from "../componentes/Header/Header";
  
const Home = () => {
  return (
    <div class="Planas">
      <Barra />
      <Header />
    </div>
    
  );
};
  
export default Home;