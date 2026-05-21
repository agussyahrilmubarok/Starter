import { Toaster } from "react-hot-toast";
import AppRoutes from "./routes";

function App() {
  return (
    <>
      <div className="container mt-5">
        <AppRoutes />
      </div>
      <Toaster position="top-right" />
    </>
  );
}

export default App;
