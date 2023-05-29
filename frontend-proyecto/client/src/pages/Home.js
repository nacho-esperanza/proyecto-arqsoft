import React from "react";
import Barra from "../componentes/Barra/Barra";

import '../App.css';

import {useParams} from "react-router-dom";
  
const Home = () => {
  return (
    <div>
      <Barra />
    </div>
    
  );
};
  
export default Home;