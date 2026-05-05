// API configuration utility
export const getApiUrl = () => {
  const apiUrl = import.meta.env.VITE_API_URL || 'http://localhost';
  const apiPort = import.meta.env.VITE_API_PORT || '8080';
  return `${apiUrl}:${apiPort}`;
};

// Helper function to make API requests
export const apiCall = async (endpoint, options = {}) => {
  const apiUrl = getApiUrl();
  const response = await fetch(`${apiUrl}${endpoint}`, options);
  return response;
};
