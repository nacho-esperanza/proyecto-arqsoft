# proyecto-arqsoft

## FRAN Y SANTI AVISO

Para cuando vayan a rendir el final, les dejo un par de tips:

- Tienen que tener en cada una de sus compus el Docker funcionando (es todo un tema y yo lo hice recien al ULTIMO de todo).
- Para ingresar como admin, Mail: admin@gmail.com Contraseña:admin
- Manejo de reservas ESTA BIEN, les recomiendo ni tocarlo.
- Si les llega a tomar el final el Franco, no le gustó NADA como esta armado el tema de los hoteles y los amenities. El quiere que cambiemos TODO esa parte y que los amenities sean una entidad distinta en donde haya una tabla pasarela con relacion muchos a muchos y se puedan agregar infinitos amenities (totalmente distinto a como lo pensamos nosotros). Cualquier cosa me preguntan pero la realidad es que hay que cambiar todas las partes del backend.
- Si tienen problemas para que aparezcan las imagenes y les saltan errores en la consola del navegador van a tener que descargarse una extension para desactivar el CORS (haganse los boludos con eso igual).
- Nosotros cargamos las imagenes con URL de internet. Pero el profe quiere que las imagenes sean cargadas desde archivos que esten en la carpeta del proyecto. Yo lo veo un poco sin sentido porque por algo lo piola es guardar las url en la base de datos. Pero weno, si quieren cambienlo y sino peleenlo como yo xd.
- El Crear Hotel (metanle un poco de css porq esta pelado) en el front tiene un bug para ingresar las estrellas (NUNCA LO PUDE ARREGLAR NO SE PORQUE), me re estresó hacer eso jajsja.
- El frontend de ingresar las imagenes al hotel no lo hice porque me dio paja xd.
- La comprobación de realizar funciones solo para ADMIN esta toda hecha en frontend. Emi me dijo que estaba correcto, pero el Franco me dijo que deberiamos realizar esa comprobación en el backend también.

Les recomiendo fijarse todas esas cosas porque la tuve que pelear mucho para aprobar.
SI ME LLEGO A ACORDAR DE ALGO MÁS LO METO ACÁ

### Bajar cambios del proyecto:

```bash
git pull origin main
```


## Comandos de GO

### Corroborar la instalación y versión de GO

```bash
go version
```

### Correr un programa GO

> Pararse siempre en la carpeta que contiene el archivo `main.go`

```bash
go run main.go
```

### Compilar y correr el programa en 2 pasos

> Pararse siempre en la carpeta que contiene el archivo `main.go`<br/>Para Linux o macOS, quitar la extensión `.exe`

```bash
go build
./nombre_del_ejecutable.exe
```

### Crear un módulo para poder usar paquetes

> Pararse siempre en la carpeta que contiene el archivo `main.go`

```bash
go mod init
go mod tidy
```

> Para todos los imports, usar el prefijo `nombre_del_modulo`, por ejemplo `import "go-api/my/package"`
>
>
 # Front

 ### Instalar dependencias

> Pararse siempre en la carpeta "frontend-vite"

```npm install
```

### Correr el front

> Pararse siempre en la carpeta "frontend-vite"

```npm run dev
```
