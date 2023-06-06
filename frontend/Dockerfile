# Utiliza una imagen base con Node.js y Next.js
FROM node:18

# Establece el directorio de trabajo en el contenedor
WORKDIR /app/frontend

# Copia los archivos de configuración del frontend
COPY package.json .
COPY package-lock.json .

# Instala las dependencias del frontend
RUN npm install

# Copia el código fuente del frontend al contenedor
COPY . .

# Genera la versión de producción del frontend
RUN npm run build

# Comando para iniciar el frontend
CMD ["npm", "start"]
