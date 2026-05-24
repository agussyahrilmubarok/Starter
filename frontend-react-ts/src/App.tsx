import { Toaster } from "react-hot-toast";
import AppRoutes from "./routes";

function App() {
  return (
    <>
      <div>
        <AppRoutes />
      </div>
      <Toaster position="top-right" />
    </>
  );
}

export default App;
