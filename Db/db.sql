-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Servidor: 127.0.0.1
-- Tiempo de generación: 27-06-2023 a las 01:18:57
-- Versión del servidor: 10.4.28-MariaDB
-- Versión de PHP: 8.2.4

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Base de datos: `go_booking_db`
--

-- --------------------------------------------------------

--
-- Estructura de tabla para la tabla `bookings`
--

CREATE TABLE `bookings` (
  `id` int(11) NOT NULL,
  `hotel_id` int(11) DEFAULT NULL,
  `user_id` int(11) DEFAULT NULL,
  `start_date` date DEFAULT NULL,
  `end_date` date DEFAULT NULL,
  `total_price` double DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Volcado de datos para la tabla `bookings`
--

INSERT INTO `bookings` (`id`, `hotel_id`, `user_id`, `start_date`, `end_date`, `total_price`) VALUES
(1, 5, 1, '2023-05-28', '2023-05-30', 300),
(2, 4, 1, '2023-05-28', '2023-05-30', 400),
(3, 4, 6, '2023-07-20', '2023-07-22', 400),
(4, 4, 6, '2023-07-20', '2023-07-22', 400),
(5, 4, 6, '2023-07-19', '2023-07-23', 800);

-- --------------------------------------------------------

--
-- Estructura de tabla para la tabla `hotels`
--

CREATE TABLE `hotels` (
  `id` int(11) NOT NULL,
  `name` varchar(50) NOT NULL,
  `adress` varchar(100) NOT NULL,
  `city` varchar(100) DEFAULT NULL,
  `stars` char(1) NOT NULL,
  `description` varchar(255) NOT NULL,
  `price` float NOT NULL,
  `parking` tinyint(1) NOT NULL,
  `pool` tinyint(1) NOT NULL,
  `wifi` tinyint(1) NOT NULL,
  `air` tinyint(1) NOT NULL,
  `gym` tinyint(1) NOT NULL,
  `spa` tinyint(1) NOT NULL,
  `rooms` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Volcado de datos para la tabla `hotels`
--

INSERT INTO `hotels` (`id`, `name`, `adress`, `city`, `stars`, `description`, `price`, `parking`, `pool`, `wifi`, `air`, `gym`, `spa`, `rooms`) VALUES
(1, 'Hotel del Viento - Islas Canarias', 'Avenida Gran Canaria 38 .  35100 Playa del Inglés España', 'Islas Canarias', '4', 'Nuestra sede de Hotel del Viento en Islas Canarias está ubicado en Playa del Inglés, una de las zonas turísticas más populares en el sur de la isla de Gran Canaria. Visita este gran hotel vacacional, garantía de unas divertidas vacaciones para todos.', 130, 1, 1, 1, 1, 1, 0, 5),
(2, 'Hotel del Viento - Benidorm', 'Av. de Filipinas, 10, 03503 Benidorm España', 'Benidorm', '5', 'Nuestra sede de Hotel del Viento en Benidorm te acerca a lo mejor de Benidorm, tratando de que disfrutes de una estancia relajante y agradable. Siendo uno de los hoteles más premium de toda España y brindando el mejor servicio del país.', 160, 1, 1, 1, 1, 0, 0, 4),
(4, 'Hotel del Viento - Caribe', 'Calle Sanchez, 94000 ', 'Santo Domingo', '4', 'Nuestra sede en el Caribe, uno de los mejores lugares para pasar tu estadia cerca de las mejores playas del mundo. El establecimiento se encuentra a 600 metros de la playa Cocolandia y a 800 metros de la playa Palenque.', 200, 1, 1, 0, 0, 0, 0, 2),
(5, 'Hotel del Viento - Milan', 'Via Cimarosa 13, 48015 Milano Marittima', 'Milan', '3', 'Nuestra sede en Italia. Se encuentra frente a la playa de Milano Marittima. Los huéspedes reciben descuentos en las instalaciones de playa cercanas. La estación de tren de Cervia está a 2 km. Nuestros clientes dicen que esta parte de Milan es su favorita.', 150, 1, 1, 0, 0, 0, 0, 1);

-- --------------------------------------------------------

--
-- Estructura de tabla para la tabla `images`
--

CREATE TABLE `images` (
  `id` int(11) NOT NULL,
  `url` varchar(255) NOT NULL,
  `hotel_id` int(11) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Volcado de datos para la tabla `images`
--

INSERT INTO `images` (`id`, `url`, `hotel_id`) VALUES
(1, 'https://media-cdn.tripadvisor.com/media/photo-s/16/1a/ea/54/hotel-presidente-4s.jpg', 2),
(2, 'https://dynamic-media-cdn.tripadvisor.com/media/photo-o/13/e9/ce/cf/hotel-presidente-4s-habitacion.jpg?w=1200&h=-1&s=1', 2),
(3, 'https://static.barcelo.com/content/dam/bhg/master/es/hoteles/spain/canarias/gran-canaria/occidental-margaritas/galeria/piscina/BMAR_POOL_06.jpg.bhgimg.stripe1000.jpg/1658834787901.jpg', 1),
(4, 'https://static.barcelo.com/content/dam/bhg/master/es/hoteles/spain/canarias/gran-canaria/occidental-margaritas/galeria/habitaciones/BMAR_ROOM_01.jpg.bhgimg.stripe1000.jpg/1658834448420.jpg', 1),
(5, 'https://cf.bstatic.com/xdata/images/hotel/max1024x768/201912067.jpg?k=8c77b0483b21335bf11980e76f1be3e9e2ab9d155d282f70eaf9c3f43272494a&o=&hp=1', 4),
(6, 'https://cf.bstatic.com/xdata/images/hotel/max1024x768/200856460.jpg?k=727b091f6b86375191fc5cb68367bffb690711bad162cedb45a7de151ac0763d&o=&hp=1', 4),
(7, 'https://cf.bstatic.com/xdata/images/hotel/max1024x768/306834049.jpg?k=d16d69f6623388e716703416bdb2a0b10754aae4c7b36c28098d8bf434fd5b56&o=&hp=1', 5),
(8, 'https://cf.bstatic.com/xdata/images/hotel/max1024x768/367950216.jpg?k=9040e8b910eda4fa6ca4dcfc74a8d80ed8df8de7ffd39d2c1675272577ba908b&o=&hp=1', 5);

-- --------------------------------------------------------

--
-- Estructura de tabla para la tabla `users`
--

CREATE TABLE `users` (
  `id` int(11) NOT NULL,
  `name` varchar(50) NOT NULL,
  `last_name` varchar(50) NOT NULL,
  `email` varchar(100) DEFAULT NULL,
  `password` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Volcado de datos para la tabla `users`
--

INSERT INTO `users` (`id`, `name`, `last_name`, `email`, `password`) VALUES
(1, 'Nacho', 'Esperanza', 'nachoy2j2@gmail.com', 'nachoo'),
(2, 'Leo', 'Miche', 'leomiche@gmail.com', '4321'),
(3, 'santi', 'nacho', 'santinacho@gmail.com', 'b59c67bf196a4758191e42f76670ceba'),
(4, 'Nacho', 'Esperanza', 'nacho@gmail.com', '68f204f3838bfee4ada868b66e6a0814'),
(6, 'Nacho', 'contraseña1602', 'nachoesperanza11@gmail.com', '68f204f3838bfee4ada868b66e6a0814'),
(7, 'Angelo', 'Capello', 'angelocapo@gmail.com', 'f70621f89830419e82ec5d4d0231f477'),
(8, 'tomas', 'ossana', 'ossanatomas@gmail.com', '81dc9bdb52d04dc20036dbd8313ed055'),
(9, 'Nachito', 'Esperanza', 'admin@gmail.com', '21232f297a57a5a743894a0e4a801fc3'),
(10, 'Cintia', 'Silvera', 'cinsilver@gmail.com', '9fc74e0f04caf7e95267dd44797cc079'),
(11, 'Pri', 'Esperanza', 'pri@gmail.com', 'c29844260862ade49405eaabe232faeb'),
(12, 'Santi', 'Master', 'santimaster@gmail.com', 'ae1d4b431ead52e5ee1788010e8ec110'),
(13, 'Fede', 'Mister', 'fedecaceres@gmail.com', '7d11810cf99c74a1f3fa22c3879ea39d'),
(14, 'Lucia', 'Chaile', 'lucichai@hotmail.com.ar', '3ba430337eb30f5fd7569451b5dfdf32'),
(15, 'Juliana', 'Marengo', 'julim@yahoo.com', '4c37dbfae76a9a48544d7248127d2d29'),
(16, 'Emi', 'Profe', 'emi@gmail.com', '12b41c761b41698d39ef68fdd9429578'),
(17, 'Fede', 'Caceres', 'fedee@gmail.com', '7d11810cf99c74a1f3fa22c3879ea39d'),
(18, 'DEmy', 'Dark', 'demydark@gmail.com', '828d3a6e7d43f3a0c85f7d9e4e432038');

--
-- Índices para tablas volcadas
--

--
-- Indices de la tabla `bookings`
--
ALTER TABLE `bookings`
  ADD PRIMARY KEY (`id`);

--
-- Indices de la tabla `hotels`
--
ALTER TABLE `hotels`
  ADD PRIMARY KEY (`id`);

--
-- Indices de la tabla `images`
--
ALTER TABLE `images`
  ADD PRIMARY KEY (`id`);

--
-- Indices de la tabla `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT de las tablas volcadas
--

--
-- AUTO_INCREMENT de la tabla `bookings`
--
ALTER TABLE `bookings`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;

--
-- AUTO_INCREMENT de la tabla `hotels`
--
ALTER TABLE `hotels`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;

--
-- AUTO_INCREMENT de la tabla `images`
--
ALTER TABLE `images`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=9;

--
-- AUTO_INCREMENT de la tabla `users`
--
ALTER TABLE `users`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=19;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
