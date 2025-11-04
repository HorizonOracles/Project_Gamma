// Animated Network Connection Background
// Canvas-based particle animation with connecting lines
import { useEffect, useRef } from "react";

interface Particle {
  x: number;
  y: number;
  vx: number;
  vy: number;
  radius: number;
  baseVx: number;
  baseVy: number;
}

interface NetworkBackgroundProps {
  className?: string;
  color?: "orange" | "gray" | "purple" | "yellow" | "gold" | "bronze" | "silver" | "sapphire" | "emerald" | "ruby" | "diamond" | "pearl" | "opal" | "stardust" | "nebula" | "supernova";
  sizeMultiplier?: number; // Multiplier for dot size (default 1.0)
}

export function NetworkBackground({ className = "", color = "orange", sizeMultiplier = 1.0 }: NetworkBackgroundProps) {
  const canvasRef = useRef<HTMLCanvasElement>(null);
  const particlesRef = useRef<Particle[]>([]);
  const animationFrameRef = useRef<number>();
  const mouseRef = useRef({ x: -1000, y: -1000 }); // Start off-screen

  useEffect(() => {
    const canvas = canvasRef.current;
    if (!canvas) return;

    const ctx = canvas.getContext("2d");
    if (!ctx) return;

    // Set canvas size with device pixel ratio for crisp rendering
    const resizeCanvas = () => {
      const dpr = window.devicePixelRatio || 1;
      const rect = canvas.getBoundingClientRect();
      
      // Set canvas internal size (accounting for DPR for crisp rendering)
      canvas.width = rect.width * dpr;
      canvas.height = rect.height * dpr;
      
      // Scale the context to match DPR (reset happens automatically when width/height change)
      ctx.scale(dpr, dpr);
    };
    resizeCanvas();
    window.addEventListener("resize", resizeCanvas);

    // Initialize particles - spread out with fewer particles
    const rect = canvas.getBoundingClientRect();
    const area = rect.width * rect.height;
    const particleCount = Math.max(
      Math.floor(area / 25000), // Much lower density for spread out look
      5 // Minimum particles even for small canvases
    );
    
    // Create particles with better initial distribution
    particlesRef.current = Array.from({ length: particleCount }, (_, i) => {
      // Use grid-based initialization for better spread, then add randomness
      const gridSize = Math.ceil(Math.sqrt(particleCount));
      const gridX = (i % gridSize) / gridSize;
      const gridY = Math.floor(i / gridSize) / gridSize;
      
      // Add larger random offset from grid position for better spread
      const x = (gridX + Math.random()) * rect.width;
      const y = (gridY + Math.random()) * rect.height;
      
      const vx = (Math.random() - 0.5) * 0.1; // Very slow, subtle movement
      const vy = (Math.random() - 0.5) * 0.1;
      return {
        x: x % rect.width, // Wrap if needed
        y: y % rect.height,
        vx,
        vy,
        baseVx: vx, // Store original velocity
        baseVy: vy, // Store original velocity
        radius: (Math.random() * 1.04 + 1.56) * sizeMultiplier, // Apply size multiplier
      };
    });

    // Animation loop
    const animate = () => {
      if (!canvas || !ctx) return;

      const rect = canvas.getBoundingClientRect();

      // Clear canvas
      ctx.clearRect(0, 0, rect.width, rect.height);

      const particles = particlesRef.current;

      // Update and draw particles
      particles.forEach((particle, i) => {
        // Mouse interaction - repel particles from cursor (more reactive)
        const mouse = mouseRef.current;
        const dx = particle.x - mouse.x;
        const dy = particle.y - mouse.y;
        const distance = Math.sqrt(dx * dx + dy * dy);
        const minDistance = 200; // Increased from 120 for larger interaction area

        if (distance < minDistance) {
          // Push particles away from mouse with stronger force
          const force = (minDistance - distance) / minDistance;
          particle.vx = particle.baseVx + (dx / distance) * force * 5; // Increased from 2 to 5
          particle.vy = particle.baseVy + (dy / distance) * force * 5;
        } else {
          // Return to base velocity
          particle.vx += (particle.baseVx - particle.vx) * 0.05;
          particle.vy += (particle.baseVy - particle.vy) * 0.05;
        }

        // Update position
        particle.x += particle.vx;
        particle.y += particle.vy;

        // Wrap around edges
        if (particle.x < 0) particle.x = rect.width;
        if (particle.x > rect.width) particle.x = 0;
        if (particle.y < 0) particle.y = rect.height;
        if (particle.y > rect.height) particle.y = 0;

        // Draw solid particle with minimal glow
        const particleSize = particle.radius; // Use consistent small size for all colors
        const colorMap: Record<string, string> = {
          gray: "rgba(140, 140, 140, 0.9)",
          purple: "rgba(140, 110, 235, 1)",
          yellow: "rgba(255, 215, 0, 1)",
          bronze: "rgba(205, 127, 50, 1)",
          silver: "rgba(192, 192, 192, 1)",
          sapphire: "rgba(15, 82, 186, 1)",
          emerald: "rgba(0, 255, 65, 1)",
          gold: "rgba(255, 215, 0, 1)",
          ruby: "rgba(220, 20, 60, 1)",
          diamond: "rgba(185, 242, 255, 1)",
          pearl: "rgba(240, 234, 214, 1)",
          opal: "rgba(168, 195, 188, 1)",
          stardust: "rgba(255, 228, 181, 1)",
          nebula: "rgba(157, 78, 221, 1)",
          supernova: "rgba(255, 107, 53, 1)",
          orange: "rgba(255, 71, 143, 0.95)"
        };
        ctx.fillStyle = colorMap[color] || colorMap.orange;
        ctx.beginPath();
        ctx.arc(particle.x, particle.y, particleSize, 0, Math.PI * 2);
        ctx.fill();

        // Draw lines to nearby particles - fewer lines with shorter range
        for (let j = i + 1; j < particles.length; j++) {
          const other = particles[j];
          const dx = particle.x - other.x;
          const dy = particle.y - other.y;
          const distance = Math.sqrt(dx * dx + dy * dy);

          const connectionRange = 120;
          if (distance < connectionRange) {
            const baseOpacity = color === "gray" ? 0.15 : 0.3;
            const opacity = (1 - distance / connectionRange) * baseOpacity;
            const lineColorMap: Record<string, string> = {
              gray: `rgba(120, 120, 120, ${opacity})`,
              purple: `rgba(140, 110, 235, ${opacity})`,
              yellow: `rgba(255, 215, 0, ${opacity})`,
              bronze: `rgba(205, 127, 50, ${opacity})`,
              gold: `rgba(255, 215, 0, ${opacity})`,
              silver: `rgba(192, 192, 192, ${opacity})`,
              sapphire: `rgba(15, 82, 186, ${opacity})`,
              emerald: `rgba(0, 255, 65, ${opacity})`,
              ruby: `rgba(220, 20, 60, ${opacity})`,
              diamond: `rgba(185, 242, 255, ${opacity})`,
              pearl: `rgba(240, 234, 214, ${opacity})`,
              opal: `rgba(168, 195, 188, ${opacity})`,
              stardust: `rgba(255, 228, 181, ${opacity})`,
              nebula: `rgba(157, 78, 221, ${opacity})`,
              supernova: `rgba(255, 107, 53, ${opacity})`,
              orange: `rgba(255, 71, 143, ${opacity})`
            };
            ctx.strokeStyle = lineColorMap[color] || lineColorMap.orange;
            ctx.lineWidth = 1.5;
            ctx.beginPath();
            ctx.moveTo(particle.x, particle.y);
            ctx.lineTo(other.x, other.y);
            ctx.stroke();
          }
        }
      });

      // Draw interactive lines from multiple nearby particles to mouse cursor
      const mouse = mouseRef.current;
      if (mouse.x > 0 && mouse.y > 0 && mouse.x < canvas.width && mouse.y < canvas.height) {
        const maxInteractionDistance = 300; // Max distance to show lines
        let hasConnection = false;
        
        particles.forEach((particle) => {
          const dx = particle.x - mouse.x;
          const dy = particle.y - mouse.y;
          const distance = Math.sqrt(dx * dx + dy * dy);
          
          // Draw line if particle is within interaction distance
          if (distance < maxInteractionDistance) {
            hasConnection = true;
            const opacity = (1 - distance / maxInteractionDistance) * 0.5;
            const mouseLineColorMap: Record<string, string> = {
              gray: `rgba(140, 140, 140, ${opacity})`,
              purple: `rgba(140, 110, 235, ${opacity})`,
              yellow: `rgba(255, 215, 0, ${opacity})`,
              bronze: `rgba(205, 127, 50, ${opacity})`,
              gold: `rgba(255, 215, 0, ${opacity})`,
              silver: `rgba(192, 192, 192, ${opacity})`,
              sapphire: `rgba(15, 82, 186, ${opacity})`,
              emerald: `rgba(0, 255, 65, ${opacity})`,
              ruby: `rgba(220, 20, 60, ${opacity})`,
              diamond: `rgba(185, 242, 255, ${opacity})`,
              pearl: `rgba(240, 234, 214, ${opacity})`,
              opal: `rgba(168, 195, 188, ${opacity})`,
              stardust: `rgba(255, 228, 181, ${opacity})`,
              nebula: `rgba(157, 78, 221, ${opacity})`,
              supernova: `rgba(255, 107, 53, ${opacity})`,
              orange: `rgba(255, 71, 143, ${opacity})`
            };
            ctx.strokeStyle = mouseLineColorMap[color] || mouseLineColorMap.orange;
            
            ctx.lineWidth = 1.5;
            ctx.beginPath();
            ctx.moveTo(particle.x, particle.y);
            ctx.lineTo(mouse.x, mouse.y);
            ctx.stroke();
          }
        });

        // Draw a small circle at mouse position if any connection exists
        if (hasConnection) {
          const mouseDotColorMap: Record<string, string> = {
            gray: "rgba(140, 140, 140, 0.7)",
            purple: "rgba(140, 110, 235, 0.7)",
            yellow: "rgba(255, 215, 0, 0.7)",
            bronze: "rgba(205, 127, 50, 0.7)",
            gold: "rgba(255, 215, 0, 0.7)",
            silver: "rgba(192, 192, 192, 0.7)",
            sapphire: "rgba(15, 82, 186, 0.7)",
            emerald: "rgba(0, 255, 65, 0.7)",
            ruby: "rgba(220, 20, 60, 0.7)",
            diamond: "rgba(185, 242, 255, 0.7)",
            pearl: "rgba(240, 234, 214, 0.7)",
            opal: "rgba(168, 195, 188, 0.7)",
            stardust: "rgba(255, 228, 181, 0.7)",
            nebula: "rgba(157, 78, 221, 0.7)",
            supernova: "rgba(255, 107, 53, 0.7)",
            orange: "rgba(255, 71, 143, 0.7)"
          };
          ctx.fillStyle = mouseDotColorMap[color] || mouseDotColorMap.orange;
          
          ctx.beginPath();
          ctx.arc(mouse.x, mouse.y, 4, 0, Math.PI * 2);
          ctx.fill();
        }
      }

      animationFrameRef.current = requestAnimationFrame(animate);
    };

    // Mouse tracking - use document level to work with pointer-events-none
    const handleMouseMove = (e: MouseEvent) => {
      const rect = canvas.getBoundingClientRect();
      const x = e.clientX - rect.left;
      const y = e.clientY - rect.top;
      
      // Only update if mouse is within canvas bounds
      if (x >= 0 && x <= rect.width && y >= 0 && y <= rect.height) {
        mouseRef.current = { x, y };
      } else {
        mouseRef.current = { x: -1000, y: -1000 };
      }
    };

    // Use document-level listeners to work through pointer-events-none
    document.addEventListener("mousemove", handleMouseMove);
    document.addEventListener("mouseleave", () => {
      mouseRef.current = { x: -1000, y: -1000 };
    });

    animate();

    return () => {
      window.removeEventListener("resize", resizeCanvas);
      document.removeEventListener("mousemove", handleMouseMove);
      if (animationFrameRef.current) {
        cancelAnimationFrame(animationFrameRef.current);
      }
    };
  }, [color, sizeMultiplier]);

  return (
    <canvas
      ref={canvasRef}
      className={className}
      style={{ display: "block", width: "100%", height: "100%" }}
    />
  );
}
