import { useState } from 'react';
import { getAll } from '../service/fetch';
export const useGetData = () => {
  const [error, setError] = useState<string | null>(null);
  const [isLoading, setIsLoading] = useState<boolean>(false);
  const [data, setData] = useState<GetData[]>([]);

  const loadData = async () => {
    try {
      setIsLoading(true);
      const response: GetData[] = await getAll();
      setData(response);
    } catch (err) {
      if (err instanceof Error) {
        setError(err.message);
      }
    } finally {
      setIsLoading(false);
    }
  };

  return { error, isLoading, data, loadData };
};
