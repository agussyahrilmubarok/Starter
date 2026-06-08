import { watchEffect } from "vue";
import { APP_NAME } from "../../constants/app";

const useDocumentTitle = (title: string, appName = APP_NAME) => {
  watchEffect(() => {
    document.title = `${title} — ${appName}`;
  });
};

export default useDocumentTitle;