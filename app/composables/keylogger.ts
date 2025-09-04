
type Options = {
  timeout?: number;
};

import { onUnmounted, readonly, ref } from 'vue';

export const useKeyLogger = (callback?: (logs:string) => void, opts?: Options) => {
  const keys = ref<string>();
  let timeoutId: ReturnType<typeof setTimeout>;
  const logKey = (event: KeyboardEvent) => {
    event.preventDefault();
    console.log(`Key pressed: ${event.key} - Code: ${keys.value}`);
    if (event.key === 'Backspace') {
      keys.value = keys.value ? keys.value.slice(0, -1) : '';
      return;
    }
    if (event.key.length > 1) {
      // Ignore special keys like Shift, Ctrl, etc.
      if (['Enter', 'Tab', 'Space'].includes(event.key)) {
        if (callback && keys.value) {
          callback(keys.value);
        }
        keys.value = ''
      }
      return;
    }
    keys.value = keys.value ? keys.value + event.key : event.key;
    if (opts?.timeout) {
      clearTimeout(timeoutId);
      timeoutId = setTimeout(() => {
        if (callback && keys.value) {
          callback(keys.value);
        }
        keys.value = '';
      }, opts.timeout);
    }
  };

  const startLogging = () => {
    window.addEventListener('keydown', logKey);
  };

  const stopLogging = () => {
    window.removeEventListener('keydown', logKey);
  };

  const reset = () => {
    keys.value = '';
  };

  onUnmounted(() => {
    stopLogging();
  });

  return {
    startLogging,
    stopLogging,
    logs: readonly(keys),
    reset,
  };
};
