import { useEffect } from "react";
import { APP_CONSTANT } from "../../utils/constant";

const useDocumentTitle = (title: string, appName = APP_CONSTANT.NAME) => {
  useEffect(() => {
    document.title = `${title} — ${appName}`;
  }, [title, appName]);
};

export default useDocumentTitle;