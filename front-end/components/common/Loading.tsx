// components/Loading.tsx

const Loading: React.FC = () => {
  return (
    <div className="flex flex-col items-center justify-center h-screen w-screen bg-gray-100">
      <p className="mb-4 text-lg font-semibold">Now Loading</p>
      <div className="flex space-x-2">
        <div className="w-3 h-3 bg-teal-500 rounded-full animate-move"></div>
        <div className="w-3 h-3 bg-teal-500 rounded-full animate-move delay-100"></div>
        <div className="w-3 h-3 bg-teal-500 rounded-full animate-move delay-200"></div>
        <div className="w-3 h-3 bg-teal-500 rounded-full animate-move delay-300"></div>
        <div className="w-3 h-3 bg-teal-500 rounded-full animate-move delay-400"></div>
      </div>
    </div>
  );
};

export default Loading;
