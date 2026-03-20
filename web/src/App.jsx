import { BrowserRouter, Routes, Route } from "react-router-dom"
import Layout from "./components/Layout"
import Home from "./page/Home"

function App() {
  return(
    <BrowserRouter>
      <Routes>
        {/* Adding element element={} */}
        <Route path="/dashboard" element={<Layout />}>
          <Route path="home" element={<Home />}></Route>
        </Route>
      </Routes>
    </BrowserRouter>
  )
}

export default App