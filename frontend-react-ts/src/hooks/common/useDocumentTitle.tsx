import { useEffect } from "react";

const useDocumentTitle = (title: string, appName = "frontend-react-ts") => {
  useEffect(() => {
    document.title = `${title} — ${appName}`;
  }, [title, appName]);
};

export default useDocumentTitle;