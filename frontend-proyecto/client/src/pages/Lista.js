import React from "react";

import '../App.css';
  
const Lista = () => {
  return (
    <div>
    hoteles
    <button onClick={() => window.location.href = 'lista/hotel'}>Este si!</button>
    <button onClick={() => window.location.href = '/#'}>Back</button>
    </div>
  );
};
  
export default Lista;