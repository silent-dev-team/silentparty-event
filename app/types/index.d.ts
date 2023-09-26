import { declare } from '../.nuxt/types/imports';
declare global {
  interface CustomerData {
    firstName: string;
    lastName: string;
    street: string;
    housenumber: string;
    zipCode: string;
    place: string;
    email: string;
  }

  interface Ticket extends CustomerData {
    sold: boolean;
    validated:  boolean;
    used: boolean;
  }
}