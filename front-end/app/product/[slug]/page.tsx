// app/blog/[slug]/page.tsx

import { notFound } from 'next/navigation';
import {sleep} from '@/app/utils/Helper';
import { Suspense } from 'react';
import Loading from '@/components/common/Loading';

interface BlogPost {
  title: string;
  content: string;
}

const getBlogPost = async (slug: string): Promise<BlogPost | null> => {
    await sleep(2000)
  return {
    title: 'hi',
    content: 'hello'
  };
};

const BlogPostPage = async ({ params }: { params: { slug: string } }) => {
  const { slug } = params;
  const post = await getBlogPost(slug);

  if (!post) {
    notFound();
  }

  return (
    <Suspense fallback={<Loading />}>
    <div>
      <h1>{post.title}</h1>
      <div>{post.content}</div>
    </div>
    </Suspense>
  );
};

export default BlogPostPage;
