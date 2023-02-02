import SystemInfo from './SystemInfo';
import DiskInfo from './DiskInfo';
import History from './History';
import Api from './Api';
import { useState, useEffect } from "react";

function App() {

  const [systemInfo, setSystemInfo] = useState(null);

  useEffect(() => {
    const updateSystemInfo = () => {
      Api.system().then(({ data }) => {
        setSystemInfo(data);
      });
    }

    const timer = setInterval(updateSystemInfo, 3000);
    return () => clearInterval(timer);
  }, []);

  return (
    <div className="app py-8 bg-slate-200 min-h-screen">
      <div className="max-w-screen-xl mx-auto">
        <SystemInfo data={systemInfo} />
        <DiskInfo data={systemInfo} />

        <History />
      </div>
    </div>
  )
}

export default App
