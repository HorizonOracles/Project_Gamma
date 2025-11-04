import loadingGif from "@assets/banner-background.gif";

export function LoadingSpinner() {
  return (
    <img 
      src={loadingGif} 
      alt="Loading..." 
      className="h-16 w-16 mx-auto mb-4"
      style={{ 
        filter: 'brightness(0) saturate(100%) invert(79%) sepia(17%) saturate(1079%) hue-rotate(358deg) brightness(95%) contrast(88%)'
      }}
    />
  );
}
