// Import the functions you need from the SDKs you need
import { initializeApp } from "firebase/app";
import { getAuth } from "firebase/auth";

// TODO: Add SDKs for Firebase products that you want to use
// https://firebase.google.com/docs/web/setup#available-libraries
// Your web app's Firebase configuration

const firebaseConfig = {
  apiKey: "AIzaSyC3I0gcRHE2w1SrA7Ok3lbJCbOYsfRclAs",
  authDomain: "focus-app-97ab7.firebaseapp.com",
  projectId: "focus-app-97ab7",
  storageBucket: "focus-app-97ab7.appspot.com",
  messagingSenderId: "706557260909",
  appId: "1:706557260909:web:695e492fc8bbe0818ffb64"
};

// Initialize Firebase
const app = initializeApp(firebaseConfig);

// Initialize Firebase Authentication and get a reference to the service
export const auth = getAuth(app);
export default app;