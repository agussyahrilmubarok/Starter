import { useEffect } from "react";
import { APP_NAME } from "../../constants/app";

const useDocumentTitle = (title: string, appName = APP_NAME) => {
  useEffect(() => {
    document.title = `${title} — ${appName}`;
  }, [title, appName]);
};

export default useDocumentTitle;
