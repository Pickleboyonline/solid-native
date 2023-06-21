type TextProps = {
  title: string;
  children: JSX.Element[];
};

export function Text({}: TextProps) {
  return <sn_text />;
}

const Bruh2 = () => {
  return <text />;
};

const Bruh = () => {
  return (
    <Text title="sdasdas">
      <Bruh2 />
    </Text>
  );
};
