import http from 'k6/http';
import { sleep } from 'k6';

export const options = {
  vus: 20,
  duration: '30s',
};

const baseURL = __ENV.BASE_URL || 'http://localhost:8082';

export default function () {
  http.get(`${baseURL}/items`);
  http.get(`${baseURL}/items?view=detailed`);
  sleep(1);
}
