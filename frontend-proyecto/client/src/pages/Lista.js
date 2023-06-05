import React from "react";

import '../App.css';
  
const Lista = () => {
  return (
    <div>
    hoteles
    <button onClick={() => window.location.href = 'lista/hotel'}>Este si!</button>
    </div>
  );
};
  
export default Lista;