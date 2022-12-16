import base64url from "base64url";
import Image, { ImageLoaderProps, ImageProps } from "next/image";
import { FC } from "react";

const IMGPROXY_URL: string = process.env.NEXT_PUBLIC_IMGPROXY_URL!;

const ImgproxyLoader = ({ src, width, quality }: ImageLoaderProps) => {
  const url = base64url(src);
  return `${IMGPROXY_URL}/sig/w:${width}/q:${quality || 80}/${url}`;
};

const ImgproxyImage: FC<ImageProps> = (props) => {
  return <Image loader={ImgproxyLoader} {...props} />;
};

export default ImgproxyImage;
